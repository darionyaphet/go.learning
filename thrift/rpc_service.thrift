namespace go rpc

service RpcService {
    string Run(1:i64 time, 2:string message, 3:map<string, string> pairs)
}