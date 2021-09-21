REM show list of devices ffmpeg -list_devices true -f dshow -i dummy
REM check resolution and format using ffmpeg -list_options true -f dshow -i video=PUT_DEVICE_NAME
go run . -rtbufsize 100M -f dshow -i video="HD User Facing" -pix_fmt yuv420p -c:v libx264 -bsf:v h264_mp4toannexb -b:v 2M -max_delay 0 -bf 0 -f h264 -