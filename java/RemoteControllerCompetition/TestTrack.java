package RemoteControllerCompetition;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class TestTrack {

    public static void race(RemoteControlCar car) {
        car.drive();
    }

    public static List<ProductionRemoteControlCar> getRankedCars(ProductionRemoteControlCar prc1,
                                                                 ProductionRemoteControlCar prc2) {
        ProductionRemoteControlCar[] cars = {prc1, prc2};
        // new ArrayList<T>(Arrays.asList(<array>)) cria uma lista mutável a
        // partir de um array
        List<ProductionRemoteControlCar> carsList = new ArrayList<ProductionRemoteControlCar>(Arrays.asList(cars));

        // Collections.sort(<list>) ordena a lista em ordem crescente a partir da
        // interface Comaparble<T> implementada pelos elementos da lista, com essa
        // lista obrigatoriamente tendo que ser mutável
        Collections.sort(carsList);

        return carsList;
    }

}
