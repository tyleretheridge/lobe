# Requirements
- Usage of FFmpeg

# FFmpeg
Extracting the embedded data streams is not currently supported. The application input
should be an already extracted binary file.

### FFmpeg Extraction Command
Sample Command to use to extract the CAMM stream embedded in a video file

```shell
<ffmpeg_path>, -y, -i, <input_video>, -map_metadata, -1, -an, -vn, -sn, \
-map, 0:d, -f, rawvideo, <path_to_output_file.bin>
```
This command is tested to work on Insta360 cameras such as the Pro2 and TITAN models. 
The largest potential source of differences for other cameras comes to identifying which stream in the video file
is the CAMM stream. If you are not sure of which stream is the CAMM stream, use FFprobe to manually inspect a sample
video file from your camera source. 


```shell
<ffprobe_path> -show_streams -of json <input_file>
```

Example stream metadata to look for:   
```
Stream #0:2[0x3](und): Data: none (camm / 0x6D6D6163), 131 kb/s
Metadata:
creation_time   : 2023-09-20T10:28:51.000000Z
handler_name    : CameraMetadataMotionHandler
```
It should be marked as a "Data" stream insteat of video or audio, and the metadata will include information or headers
that reference "camm" or "CameraMetadataMotionHandler"  
The "0:d" flag in the extraction command references the camm stream by it's type and position.   
In my case, the CAMM stream is the 0th Data stream, so it is mapped as "0:d"


### FFmpeg Resource Links
- [Github Gist dictionary of FFmpeg flag/options](https://gist.github.com/tayvano/6e2d456a9897f55025e25035478a3a50)
- [Official FFmpeg Documentation](https://ffmpeg.org/ffmpeg.html)