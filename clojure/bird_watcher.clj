(ns bird-watcher)

(def last-week
  [0 2 5 3 7 8 4])

(defn today
  [birds]
  ;; (last <col>) retorna o último elemento da coleção
  (last birds))

(defn inc-bird
  [birds]
  ;; (assoc <col> <indice> <valor>) associa (atualiza) um valor ao indicie da coleção
  (assoc birds (- (count birds) 1) (+ (today birds) 1)))

(defn day-without-birds?
  [birds]
  ;; (boolean <valor>) converte o valor para booleano
  ;;
  ;; #(<...> %) cria uma função lambda com o argumento %, equivalente a (fn [<var>] (...))
  (boolean (some #(= 0 %) birds)))

(defn n-days-count
  [birds n]
  ;; (reduce <fn> <col>) aplica a função <fn> para os items da coleção, indo de 2
  ;; em 2 items
  ;;
  ;; (take <n> <col>) obtém os primeiros n elementos da coleção
  (reduce #(+ %1 %2) (take n birds)))

(defn busy-days
  [birds]
  ;; (filter <predicado> <col>) filtra uma coleção retornando apenas os items
  ;; cujo predicado for verdadeiro
  (count (filter #(>= % 5) birds)))

(defn odd-week?
  [birds]
  (= birds [1 0 1 0 1 0 1]))
