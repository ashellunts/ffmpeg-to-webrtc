# ffmpeg-to-webrtc

ffmpeg-to-webrtc demonstrates how to send video from ffmpeg to your browser using [pion](https://github.com/pion/webrtc).

## How to run it

### Open example web page
[jsfiddle.net](https://jsfiddle.net/wjo4e9c7/1/)

### Copy browser's SDP
In the jsfiddle the top textarea is your browser's SDP, copy that to clipboard.

#### Windows
1. `cd src`
1. Paste the SDP into a file `src/SDP.txt`.
2. Make sure ffmpeg in your PATH and golang is installed.
3. Run `go run . <ffmpeg command line options> - < SDP.txt`
4. Note dash after ffmpeg options. It makes ffmpeg to write output to stdout. The app will read h264 stream from ffmpeg stdout.
5. ffmpeg output format should be h264. Browsers don't support all h264 profiles so it may not always work. Here is an example of format that works: `-pix_fmt yuv420p -c:v libx264 -bsf:v h264_mp4toannexb -b:v 2M -max_delay 0 -bf 0 -f h264`.

### Put SDP from ffmpeg-to-webrtc into your browser
When you see SDP in base64 format printed it means that SDP is already in clipboard. So you can go to jsfiddle page and paste that into Application SDP text area.

### Hit 'Start Session' in jsfiddle
A video should start playing in your browser below the input boxes.

## Examples (windows)
### Share camera stream
```go run . -rtbufsize 100M -f dshow -i video="PUT_DEVICE_NAME" -pix_fmt yuv420p -c:v libx264 -bsf:v h264_mp4toannexb -b:v 2M -max_delay 0 -bf 0 -f h264 - < SDP```. 
There is a delay of several seconds. Should be possible to fix it with better ffmpeg configuration.

To check list of devices: `ffmpeg -list_devices true -f dshow -i dummy`.  
It is possible also to set a resolution and a format, for example `-pixel_format yuyv422 -s 640x480`.
Possible formats: `ffmpeg -list_options true -f dshow -i video=PUT_DEVICE_NAME`.
### Share screen or window
See `.bat` files in src folder

## Linux, macOS

Should work on other operating systems, though I haven't tried.
