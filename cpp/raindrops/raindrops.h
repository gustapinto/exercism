#if !defined(RAINDROPS_H)
#define RAINDROPS_H
#include <string>

namespace raindrops {
    std::string convert(int n) {
        std::string sound;

        if (n % 3 == 0) {
            sound.append("Pling");
        }

        if (n % 5 == 0) {
            sound.append("Plang");
        }

        if (n % 7 == 0) {
            sound.append("Plong");
        }

        if (sound.size() == 0) {
            return std::to_string(n);
        }

        return sound;
    }
}  // namespace raindrops

#endif // RAINDROPS_H