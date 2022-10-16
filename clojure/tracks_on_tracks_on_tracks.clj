(ns tracks-on-tracks-on-tracks)

(defn new-list
  "Creates an empty list of languages to practice."
  []
  ;; Cria uma nova lista vazia usando '()
  '())

(defn add-language
  "Adds a language to the list."
  [lang-list lang]
  ;; (cons <item> <coleção>) concatena o item a uma coleção, retornando uma
  ;; nova coleção contendo os elemntos originais e o elemento concatenado no
  ;; começo
  (cons lang lang-list))

(defn first-language
  "Returns the first language on the list."
  [lang-list]
  ;; (first <coleção>) retorna o primeiro item da coleção
  (first lang-list))

(defn remove-language
  "Removes the first language added to the list."
  [lang-list]
  ;; (rest <coleção>) retorna uma coleção com todos os elementos da coleção original
  ;; exceto o primeiro
  (rest lang-list))

(defn count-languages
  "Returns the total number of languages on the list."
  [lang-list]
  ;; (count <coleção>) retorna a quantidade de elementos em uma coleção
  (count lang-list))

(defn learning-list
  "Creates an empty list, adds Clojure and Lisp, removes Lisp, adds
  Java and JavaScript, then finally returns a count of the total number
  of languages."
  []
  ;; (-> <fn 1> <fn 2>) cria um pipe, com os resultados das funções sendo
  ;; passados para a próxima função como o primeiro argumento
  (->
    (new-list)
    (add-language "Clojure")
    (add-language "Lisp")
    (remove-language)
    (add-language "Java")
    (add-language "JavaScript")
    (count-languages)))
