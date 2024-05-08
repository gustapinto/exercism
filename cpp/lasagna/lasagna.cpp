// const <T> <var> -> declara uma constante
const int TOTAL_OVEN_TIME = 40;
const int MINUTES_PER_LAYER = 2;

int ovenTime() {
    return TOTAL_OVEN_TIME;
}

int remainingOvenTime(int actualMinutesInOven) {
    return ovenTime() - actualMinutesInOven;
}

int preparationTime(int numberOfLayers) {
    return numberOfLayers * MINUTES_PER_LAYER;
}

int elapsedTime(int numberOfLayers, int actualMinutesInOven) {
    return preparationTime(numberOfLayers) + actualMinutesInOven;
}
