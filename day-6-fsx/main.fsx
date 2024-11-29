let races = [(48.0, 390.0); (98.0, 1103.0); (90.0, 1112.0); (83.0, 1360.0)]
let fatRace = [(48989083.0, 390110311121360.0)]

let calc_boundary (n, y) = n / 2.0 - sqrt (n * n/4.0 - y), n / 2.0 + sqrt (n * n/4.0 - y)
// let calc_upper_boundary (n, y) = n / 2 + (y - n/4) ** 0.5

// let lower_bounds = List.map calc_boundary races
let exampleRaces = [(7.0, 9.0); (15.0, 40.0); (30.0, 200.0);]
let exampleFatRace = [(71530.0, 940200.0)]
// let (a, b) = calc_boundary (30.0, 200.0)
// // printfn "%A or %A" (a + 1.0) (b - 1.0)  //(ceil a) (floor b)
// printfn "%A" (b - a - 1.0)

let bounds = List.map calc_boundary exampleFatRace

let calc_span (x, y) = (floor y - ceil x - 1.0) + (ceil ((ceil x) - x)) * 1.0 + (ceil (y - (floor y))) * 1.0

let res_orint (a, b)= printfn "%A" (b - a - 1.0)
let ppp x = printfn "%A" x
let numberOfImprovingPossibilities = List.map calc_span bounds
List.map ppp numberOfImprovingPossibilities
printfn "Margin of error: %A" (List.fold (*) 1.0 numberOfImprovingPossibilities)
// printfn "%A" exampleBounds