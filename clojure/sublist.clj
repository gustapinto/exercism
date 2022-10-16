(ns sublist)

(defn sublist?
  "Retorna se a primeira lista é uma sublista da segunda"
  [a b]
  ;; #{<lista>} cria um set não ordenado a partir da lista passada
  ;;
  ;; (some <coleção> <predicado>) retorna o primeiro valor lógico para o predicado
  ;; passado
  ;;
  ;; (partition <n> <step> <coleção>) retorna uma lista de listas com <n>
  ;; elementos cada, divididos em blocas de <step> elementos
  (some #{a} (partition (count a) 1 b)))

(defn classify
  [list1 list2]
  (cond (= list1 list2) :equal
        (sublist? list1 list2) :sublist
        (sublist? list2 list1) :superlist
        :else :unequal))
