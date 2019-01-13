# composal

Communal composition tool. Get it?

## install

I don't really know. You have to get your go environment setup first. Then run

`go get github.com/colevscode/composal`

I think.

## build

In the root folder just run

`go build`

and you'll have the `composal` binary ready to go

## use

Start the composal server

`./composal -port 8088 -prefix examples/samples`

Then send requests to render some songs

`curl --data @examples/song1.json http://localhost:8088/render > song.wav`

now play it

`play song.wav`

nice.
