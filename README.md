# resolveip

```bash


customIP := reqIF.ResolveIP
dialer := &net.Dialer{
    Timeout:   30 * time.Second,
    KeepAlive: 30 * time.Second,
}
fn := func(ctx context.Context, network, addr string) (net.Conn, error) {
        if addr == reqIF.Host {
        customIP = customIP + ":443"
        zap.L().Debug("Resolve IP dialer",
        zap.Any("original", addr),
        zap.Any("modified", customIP),
    )
    addr = customIP
}


return dialer.DialContext(ctx, network, addr)



//client := req.C() // Use C() to create a client.
client.SetDial(fn)


If monitor by wireShark

Check connection with Wire(wlp4)  / Cable (eth01) for wireShark
Disable all IP6



If wanna to test the “Failure” case. Must change IP front 3 digit

XXX.XXX.XXX.XXX

[static].[static].[range].[1~255]




```
