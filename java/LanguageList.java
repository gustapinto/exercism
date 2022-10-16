import java.util.ArrayList;
import java.util.List;

public class LanguageList {

    private final List<String> languages = new ArrayList<>();

    public boolean isEmpty() {
        // <List>.isEmpty() retorna se a lista está vazia
        return this.languages.isEmpty();
    }

    public void addLanguage(String language) {
        // <List>.add(<item>) adiciona um novo elemento ao final da lista
        this.languages.add(language);
    }

    public void removeLanguage(String language) {
        // <List>.remove(<item>) remove a primeira ocorrência de um item da lista
        // se ele existir
        this.languages.remove(language);
    }

    public String firstLanguage() {
        // <List>.get(<indice>) retorna o elemento presente no indice indicado
        return this.languages.get(0);
    }

    public int count() {
        // <List>.size() retorna a quantidade de elementos na lista
        return this.languages.size();
    }

    public boolean containsLanguage(String language) {
        // <List>.contains(<item>) verifica se um elemento existe na lista
        return this.languages.contains(language);
    }

    public boolean isExciting() {
        return this.containsLanguage("Java") || this.containsLanguage("Kotlin");
    }

}
