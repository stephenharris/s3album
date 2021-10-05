resource "aws_lambda_layer_version" "ffmpeg_layer" {
  filename   = "ffmpeg.zip"
  layer_name = "ffmpeg_layer"
  compatible_runtimes = ["python3.9"]
  depends_on = [
    null_resource.build_ffmpeg_layer
  ]
}

resource "null_resource" "build_ffmpeg_layer" {
 provisioner "local-exec" {
    command = "/bin/bash scripts/build-ffmpeg-layer.sh"
  }
}
