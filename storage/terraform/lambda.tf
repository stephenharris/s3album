resource "aws_lambda_function" "test_lambda" {
  filename      = "thumbnail_generator.zip"
  function_name = "ThumbnailGenerator"
  role          = aws_iam_role.thumbnail_generator_execution_role.arn
  handler       = "thumbnail_generator.lambda_handler"

  source_code_hash = filebase64sha256("thumbnail_generator.zip")

  layers = [aws_lambda_layer_version.ffmpeg_layer.arn]

  runtime = "python3.9"

  memory_size                    = 512
  timeout                    = 60

  environment {
    variables = {
      foo = "bar"
    }
  }
}

resource "aws_iam_role" "thumbnail_generator_execution_role" {
  name = "ThumbnailGeneratorExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.thumbnail_generator_execution_assume_role_policy.json
}

data "aws_iam_policy_document" "thumbnail_generator_execution_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "allow_s3_readwrite" {
  policy = data.template_file.allow_s3_readwrite.rendered
}

data "template_file" "allow_s3_readwrite" {
  template = "${file("thumbnail_generator_policy.json.tpl")}"

  vars = {
    src_bucket = "${aws_s3_bucket.photos.arn}"
    dest_bucket = "${aws_s3_bucket.photos.arn}"
  }
}

resource "aws_iam_policy_attachment" "thumbnail_generator_allow_s3_readwrite" {
  name       = "ThumbnailGeneratorExecutionRoleAllowS3"
  roles      = [aws_iam_role.thumbnail_generator_execution_role.name]
  policy_arn = aws_iam_policy.allow_s3_readwrite.arn
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromS3Event"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.photos.arn
  source_account =  data.aws_caller_identity.current.account_id
}


resource "aws_s3_bucket_notification" "file_uploaded_notification" {
  bucket = aws_s3_bucket.photos.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.test_lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = var.src_object_prefix
  }
}