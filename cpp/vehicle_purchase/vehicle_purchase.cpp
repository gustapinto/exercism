#include "vehicle_purchase.h"
#include <set>
#include <string>
#include <iostream>

const std::set<std::string> VEHICLES_THAT_REQUIRE_LICENSE{"car", "truck"};

namespace vehicle_purchase {
    bool needs_license(std::string kind){
        // <std::set<T>>.find(<T>) -> Procura um elemento <T> dentro do <std::set<T>>
        // <std::set<T>>.end() -> Retorna um iterador para o fim do <std::set<T>>
        return VEHICLES_THAT_REQUIRE_LICENSE.find(kind) != VEHICLES_THAT_REQUIRE_LICENSE.end();
    }

    std::string choose_vehicle(std::string option1, std::string option2) {
        std::string choosen = (option1 < option2) ? option1 : option2;

        return choosen + " is clearly the better choice.";
    }

    double calculate_resell_price(double original_price, double age) {
        if (age < 3) {
            return original_price * 0.8;
        }

        if (age >= 3 && age < 10) {
            return original_price * 0.7;
        }

        return original_price * 0.5;
    }

}  // namespace vehicle_purchase
