# Coliseu CLI

Coliseu is a command line video downloader and audio extractor. For now, the client is capable of fetching YouTube video metadata and downloading videos.

#### Usage

```
Coliseu [global options] command [command options] [arguments...]
```

**Examples:**

```
coliseu youtube -d https://youtu.be/pZ5576Pags4
```

or

```
coliseu youtube -d pZ5576Pags4
```

```
YouTube
Argument: pZ5576Pags4
Argument is video id
Video: pZ5576Pags4
Metadata:
 - title: The World's Smallest Dog: Tiny Dog Terrier
 - length: 0.9833333333333333 min
 - format:
     0 - hd720 video/mp4; codecs="avc1.64001F, mp4a.40.2"
     1 - medium video/webm; codecs="vp8.0, vorbis"
     2 - medium video/mp4; codecs="avc1.42001E, mp4a.40.2"
     3 - small video/x-flv
     4 - small video/3gpp; codecs="mp4v.20.3, mp4a.40.2"
     5 - small video/3gpp; codecs="mp4v.20.3, mp4a.40.2"
     6 - Cancel
Select format to download : 4
1.58 MB / 1.58 MB [================================================] 100.00 % 12s
Done
```

#### Commands

```
youtube, y	YouTube downloader
  --download, -d <url or video id>   Download video

help, h	Shows a list of commands or help for one command
```

**Global options**

```
--help, -h		show help
--version, -v	print the version
```

Install or upgrade
----

Run the following commands (in a bash shell):

```bash
curl -L https://github.com/ricardopereira/coliseu/releases/download/v0.2/coliseu-x86_64 > /usr/local/bin/coliseu
```

Then:

```bash
chmod +x /usr/local/bin/coliseu
```

Author
----

Ricardo Pereira
