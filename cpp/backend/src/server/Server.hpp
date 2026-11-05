
#pragma once
#include <iostream>
#include <string>


class Server {
public:
    Server() = default;
    ~Server() = default;

protected:
    int main(const std::vector<std::string>& args) override;
}