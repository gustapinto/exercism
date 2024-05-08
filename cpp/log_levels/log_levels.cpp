#include <algorithm>
#include <string>

const std::string TRIM_CHARS{" \n\r\t"};
const std::string LOG_DELIMITER{":"};

namespace utils {
    // &<var> -> Passa um ponteiro para a variável
    std::string trim(const std::string input) {
        std::string s{input};

        // <std::string>.erase(<pos>) -> Remove todos os elementos da string até a posição
        // <std::string>.find_first_not_of(<chars>) -> Encontra o primeiro elemento que não esteja nos chars
        s.erase(0, s.find_first_not_of(TRIM_CHARS));
        // <std::string>.find_last_not_of(<chars>) -> Encontra o último elemento que não esteja nos chars
        s.erase(s.find_last_not_of(TRIM_CHARS)+1);

        return s;
    }

    std::string upper(const std::string input) {
        std::string s{input};

        // <std::string>.begin() -> Retorna um iterador para o início da string
        // <std::string>.end() -> Retorna um iterador para o final da string
        // std::transform() -> Aplica uma função para um conjunto de elementos, salvando o resultado em um conjunto de destino
        // [](<var>) {} -> Inicia uma expressão lambda
        std::transform(s.begin(), s.end(), s.begin(), [](char c) {
            // std::toupper(<char>) -> Converte o char para maiusculo
            return std::toupper(c);
        });

        return s;
    }
}

namespace log_line {

    std::string message(std::string line) {
        // <std::string>.find_first_of(<std::string>) -> Retorna a posição da primeira ocorrência da string buscada
        int pos = line.find_first_of(LOG_DELIMITER);
        // <std::string>.substr(<pos1>, <pos2>) -> Retorna uma substring entre as duas posições passadas
        std::string message = line.substr(pos+1, line.size());

        return utils::trim(message);
    }

    std::string log_level(std::string line) {
        int pos = line.find_first_of(LOG_DELIMITER);
        std::string level = line.substr(0, pos);

        // <std::string>.replace(<pos>, <qtde>, <new>) -> Substitui a quantidade <qtde> de caracteres da string a partir da posição <pos> com a string <new>
        level.replace(level.find_first_of("["), 1, "");
        level.replace(level.find_first_of("]"), 1, "");

        return utils::upper(utils::trim(level));
    }

    std::string reformat(std::string line) {
        return message(line) + " (" + log_level(line) + ")";
    }
}
