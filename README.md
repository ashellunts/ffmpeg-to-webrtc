# ffmpeg-to-webrtc

ffmpeg-to-webrtc demonstrates how to send video from ffmpeg to your browser using [pion](https://github.com/pion/webrtc).

This example has the same structure as [play-from-disk-h264](https://github.com/pion/example-webrtc-applications/blob/master/play-from-disk-h264) but instead of reading from a file it reads H264 stream from ffmpeg stdout pipe.

## Instructions

### Open example web page
[jsfiddle.net](https://jsfiddle.net/wjo4e9c7/1/)

### Copy browser's SDP
In the jsfiddle the top textarea is your browser's SDP, copy that to clipboard.

#### Windows
1. `cd src`
1. Paste the SDP into a file `src/SDP.txt`.
2. Make sure ffmpeg in your PATH and golang is installed.
3. Run `go run . -i output.ogg -vn -c:a libopus -page_duration 20000 -f ogg pipe:1 < sdp.txt`
4. Note dash after ffmpeg options. It makes ffmpeg to write output to stdout. The app will read h264 stream from ffmpeg stdout.
5. ffmpeg output format should be h264. Browsers don't support all h264 profiles so it may not always work. Here is an example of format that works: `-pix_fmt yuv420p -c:v libx264 -bsf:v h264_mp4toannexb -b:v 2M -max_delay 0 -bf 0 -f h264`.

### Put SDP from ffmpeg-to-webrtc into your browser
When you see SDP in base64 format printed it means that SDP is already in clipboard. So you can go to jsfiddle page and paste that into Application SDP text area.

### Hit 'Start Session' in jsfiddle
A video should start playing in your browser below the input boxes.