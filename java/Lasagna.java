public class Lasagna {

    /* OBS
     *
     * Quando declaramos métodos sem nenhuma keyword de visibilidade (public,
     * private, protected) significa que esse método só pode ser acessado pelas
     * classes dentro do pacote atual
    */
    int expectedMinutesInOven() {
        return 40;
    }

    int remainingMinutesInOven(int minutesAlreadyInOven) {
        return this.expectedMinutesInOven() - minutesAlreadyInOven;
    }

    int preparationTimeInMinutes(int layers) {
        return layers * 2;
    }

    int totalTimeInMinutes(int layers, int minutesAlreadyInOven) {
        return minutesAlreadyInOven + preparationTimeInMinutes(layers);
    }

}
