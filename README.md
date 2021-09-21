# ffmpeg-to-webrtc

ffmpeg-to-webrtc demonstrates how to send video from ffmpeg to your browser.

This example has the same structure as [play-from-disk-h264](https://github.com/pion/example-webrtc-applications/blob/master/play-from-disk-h264) but instead of reading from a file it reads H264 stream from ffmpeg stdout pipe.

## Instructions

### Open example page
[jsfiddle.net](https://jsfiddle.net/9s10amwL/) you should see two text-areas and a 'Start Session' button

### Copy browser's SessionDescription
In the jsfiddle the top textarea is your browser's SDP, copy that and:

#### Windows
1. Paste the SessionDescription into a file `SDP`.
2. Run `go run . <ffmpeg command line options> - < SDP`
3. Note dash after ffmpeg options. It makes ffmpeg to write output to stdout. Example app will read stream from ffmpeg stdout.

### Input SessionDescription from ffmpeg-to-webrtc into your browser
When you see SDP in base64 format printed it means that SDP is already in copy buffer. So you can go to jsfiddle page and paste that into second text area

### Hit 'Start Session' in jsfiddle, enjoy your video!
A video should start playing in your browser below the input boxes.

## Examples
See .bat files in src folder (windows only)