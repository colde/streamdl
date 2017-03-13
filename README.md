# streamdl

[![Build Status](https://travis-ci.org/colde/streamdl.svg?branch=master)](https://travis-ci.org/colde/streamdl)

## Usage

Simply call the binary with a "max time" and a Smooth Streaming manifest url

    ./streamdl-macos -d 200ms -u http://wams.edgesuite.net/media/SintelTrailer_MP4_from_WAME/sintel_trailer-1080p.ism/Manifest

This will print any request that takes longer than the allotted time, or otherwise fails.
The fetcthing will happen with 2 simultaneous workers to simulate an actual client behaviour
(although it will fetch as fast as possible)

## Implementation details

* Only links ending in /Manifest is supported.
* 302 redirects on the Manifest is supported.