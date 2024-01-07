# What is Lobe? 

Lobe is a command line tool to decode data streams embedded in videos that conform to the CAMM Specification maintained and detailed by Google.

As a result, this tool can be used to parse data such as: 
- GPS 
- Gyroscope data
- Accelerometer data
- 3/6 DoF position

# CAMM (Camera Motion Metadata)
From the [official Google spec documentation](https://developers.google.com/streetview/publish/camm-spec):   
CAMM is "a specification that allows MP4 files to embed metadata about camera motion during video capture. 
Devices that capture video typically have sensors that can provide additional information about capture."

# Installation


# How To Use

Lobe currently accepts a binary data stream that has been extracted from a video source.
This tool does not currently have built in extraction of the data from video sources, and as such the user will have to 
use the tool in conjunction with another tool to perform this task. 

FFmpeg & FFprobe are recommended tools as the suite is the gold standard in video processing.



# Releases
None Currently Available


# Tasks
Implement Writers as a separate package, create boundary between data and output 




# Containers
Containers are configured to store specific types of packets
Should take in an options struct that decides what packets are being stored


# CSV Writer

### Options
- Check if container has one or more packet types 
- If container has multiple packet types
  - Single File (records become wrapped)
  - Multiple Files (each packet type is sent to a separate file and name is generated)

- Single Packet in container
    - Packets are written to the csv file with fields as csv headers

add a method on interface that returns stored packet types
this will be available when the container is configured in the cli from flags


If container is configured to store 1 type of packet
csv  
header1, header2, header3  
rows of values for packet type 6  

If container is configured to store 1 type of packet
packettype, headers, values  
rows of the above  


/path/to/csv/csv-typezero.csv
/path/to/csv/csv-typeone.csv
/path/to/csv/csv-typetwo.csv
/path/to/csv/csv-typethree.csv
