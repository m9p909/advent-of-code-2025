(require '[clojure.string :as str])

(print "hello world")

(defn read-from-file []
  (slurp "./input.txt"))

(def example-data {:ranges [[]]
                   :values []})
(defn parse-range [strrange]
  (map #(read-string %)
       (str/split strrange #"-")))

(defn parse-ranges [ranges]
  (map parse-range ranges))

(defn file-to-data [strdata]
  (let [lines (str/split strdata #"\n\n")
        ranges (parse-ranges (str/split (get lines 0) #"\n"))
        times (map #(read-string %) (str/split (get lines 1) #"\n"))]
    {:lines lines
     :ranges ranges
     :times times}))

;; Are the ranges in order? no
;; keep an array of start and end points of the ranges, then we can quickly check
;; or just search all ranges
;; easier to just search all ranges
;; part 2 is keep an array of start and end point of the ranges then calculate all the ranges

(defn solve-problem [data]
  (let [ranges (:ranges data)
        times (:times data)]
    (count
     (filter (fn [time]
               (some (fn [[start end]]
                       (<= start time end))
                     ranges))
             times))))

(defn merge-ranges-reducer [merged rang]
  (if (empty? merged)
    (conj merged rang)
    (let [[start end] rang
          [cstart cend]  (first merged)]
      (if (>= start cend)
        (conj merged rang)
        (conj (rest merged)  [cstart (max end cend)])))))

;; idk why this is failing
(merge-ranges-reducer '() '(21 30))
(conj '(5,8,9) 20)
(butlast '(5,8,9,20))
(let [[start end] '(10,20)]
  (print start end))

;; merges the ranges into a list of ranges with no overlap
(defn merge-ranges [ranges]
  (let [sorted (sort-by #(first %) ranges)]
    (reduce merge-ranges-reducer
            '()
            sorted)))

(defn solve-problem2 [data]
  (let [ranges (merge-ranges (:ranges data))]
    (reduce (fn [prev curr]
              (+ prev (inc (- (last curr) (first curr)))))
            0, ranges)))

(def test-data '((3 5) (10 14) (16 20) (12 18)))
(merge-ranges test-data)

(solve-problem2 (file-to-data (read-from-file))) ; 344306344403189



