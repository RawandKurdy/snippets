#include <boost/network/protocol/http/server.hpp>
#include <iostream>

namespace http = boost::network::http;

struct greet;
typedef http::server<greet> server;

struct greet {
    void operator()(server::request const &request,
    server::connection_ptr connection) {
        std::ostringstream data;
        data << "Hello World!";
        connection->set_status(server::connection::ok);
        connection->write(data.str());
    }
};

int main(int argc, char *argv[]) {
    std::cerr << "listening at http://localhost:8080" << std::endl;

    greet handler;
    server::options options(handler);
    server server_(options.address("0.0.0.0").port("8080"));
    server_.run();

    return 0;
}