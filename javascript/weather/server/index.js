const grpc = require("@grpc/grpc-js");
const { WeatherServiceService } = require("../proto/weather_grpc_pb");
const serviceImpl = require("./service_impl");

const addr = "localhost:50051";

function cleanup(server) {
    console.log("Cleanup...");
    if (server) {
        server.forceShutdown();
    }
}

function main() {
    const server = new grpc.Server();
    const creds = grpc.ServerCredentials.createInsecure();

    process.on("SIGINT", () => {
        console.log("Caught interrupt signal!");
        cleanup(server);
    })

    server.addService(WeatherServiceService, serviceImpl);

    server.bindAsync(addr, creds, (err, _) => {
        if (err) {
            return cleanup(server);
        }
        server.start();
    })
    console.log(`Server listening on port: ${addr}`);
};
main();