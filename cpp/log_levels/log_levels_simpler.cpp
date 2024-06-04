#include <algorithm>
#include <string>

const std::string TRIM_CHARS{" \n\r\t"};
const std::string LOG_DELIMITER{":"};

namespace utils {
    std::string trim(const std::string input) {
        std::string s{input};

        s.erase(0, s.find_first_not_of(TRIM_CHARS));
        s.erase(s.find_last_not_of(TRIM_CHARS)+1);

        return s;
    }

    std::string upper(const std::string input) {
        std::string s;

        for (const char &c : input) {
            // <string>.push_back(<char>) -> Add a new char at the end of the string
            s.push_back(std::toupper(c));
        }

        return s;
    }
}

namespace log_line {

    std::string message(std::string line) {
        int pos = line.find_first_of(LOG_DELIMITER);
        std::string message = line.substr(pos+1, line.size());

        return utils::trim(message);
    }

    std::string log_level(std::string line) {
        int pos = line.find_first_of(LOG_DELIMITER);
        std::string level = line.substr(0, pos);

        level.replace(level.find_first_of("["), 1, "");
        level.replace(level.find_first_of("]"), 1, "");

        return utils::upper(utils::trim(level));
    }

    std::string reformat(std::string line) {
        return message(line) + " (" + log_level(line) + ")";
    }
}
