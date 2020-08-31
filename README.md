我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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
