# 簡易WebSocket Proxyサーバー ayameproxy

## これは何?

[WebRTCシグナリングサーバーAyame](https://github.com/OpenAyame/ayame)のシグナリングURLをproxyするために使うサーバーです。

完全にテスト用です。この文書を読んで、テスト用という意味が理解できない方のご利用はお控えください。

## 何をテストするのか

ローカルネットワーク環境において、[ayame](https://github.com/OpenAyame/ayame)と[momo](https://github.com/shiguredo/momo)間などでP2PのWebRTC通信をテストするために使います。私はmomoをカスタマイズしており、momo自体の動作確認に使用しています。ayameおよびmomoは、[株式会社時雨堂](https://shiguredo.jp)さんのプロダクトです。

## なぜ作ったのか

ayameのデモ画面を利用して、ayameが動作する機械に接続されているカメラやマイクを利用するには、ayame自体をlocalhost(127.0.0.1)にバインドする必要があります。ayameのデフォルトの設定ではlocalhost:3000をリスンします。
そうすると、シグナリングURLは `ws://localhost:3000/signaling` となりますが、このURLはネットワーク越しの別の機械からは到達不能です。到達可能にするには、[ayameのUSE.md](https://github.com/OpenAyame/ayame/blob/develop/doc/USE.md)にも記述があるように、一般にはngrokなどのフォワーダ(proxy)を利用してシグナリングURLを外部公開します。ですが、ngrokをセットアップしたりTLS化したりせずに、さっとmomoとの通信テストをしたい。そんな目的のために作りました。

## 何をしているか

WebSocket通信をayameのデフォルトであるlocalhost:3000にproxyします。

## 使い方

```
% ./ayameproxy --help
Usage of ./ayameproxy:
        ./ayameproxy [-l [host]:port]
        
Options:
  -l string
        listen host and port (default ":13000")
```

デフォルトでは、13000番ポートをリスンします。

例えば、`192.168.0.1`のIPアドレスの機械で動かすと、シグナリングURLは`ws://192.168.0.1:13000/signaling`となります。

## ビルド方法

ayameを作る。

```
% git clone https://github.com/OpenAyame/ayame.git
% cd ayame
% go build
```

ayameproxyを作る。

```
% git clone https://github.com/msnoigrs/ayameproxy.git
% cd ayameproxy
% go build
```

## テスト方法

ayameとayameproxyを起動したら、別の機械からmomoなどでシグナリングURLとroom IDを指定して接続して確認。

## ライセンス

ライセンスは、NYSLです。
