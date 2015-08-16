Coliseu CLI
----

Coliseu is a command line video downloader and audio extractor. For now, the client is capable of fetching YouTube video metadata and downloading videos.

####Usage

```
Coliseu [global options] command [command options] [arguments...]
```

**Examples:**

```
coliseu youtube -d https://youtu.be/pZ5576Pags4
```

```
coliseu youtube -d pZ5576Pags4
```

####Commands

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
