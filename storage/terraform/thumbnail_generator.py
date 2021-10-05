import boto3
import os
import sys
import uuid
from urllib.parse import unquote_plus
from PIL import Image
import PIL.Image

s3_client = boto3.client('s3')

def resize_image(image_path, resized_path):
  with Image.open(image_path) as image:
      image.thumbnail(tuple(min(x / 4, 400) for x in image.size))
      image.save(resized_path)

def video_thumb(video, upload_path):
  ffpg = "/opt/bin/ffmpeg"
  ffmpeg_cmd = ffpg + " -i \"" + video + "\" -ss 00:00:00 -vframes 1 " + upload_path
  os.system(ffmpeg_cmd)

def is_video(path):
  return (
    path.lower().endswith(".avi")
    or path.lower().endswith(".3gp") 
    or path.lower().endswith(".mp4")
  )

def lambda_handler(event, context):
  for record in event['Records']:
      bucket = record['s3']['bucket']['name']
      key = unquote_plus(record['s3']['object']['key'])

      if not key.lower().startswith("thumbnail"):

        tmpkey = key.replace('/', '')
        download_path = '/tmp/{}{}'.format(uuid.uuid4(), tmpkey)
        upload_path = '/tmp/thumbnail-{}'.format(tmpkey)
        print(key)
        s3_client.download_file(bucket, key, download_path)

        if is_video(download_path):
          upload_path = upload_path + ".jpg"
          key = key + ".jpg"
          video_thumb(download_path, upload_path)
        else:
          resize_image(download_path, upload_path)

        s3_client.upload_file(upload_path, bucket, "thumbnail/" + key)
        os.remove(upload_path) 
        os.remove(download_path) 
      else:
        print("ignore " + key)
