let statesNeeded = new Set([ "mt", "wa", "or", "id", "nv", "ut", "ca", "az" ])

const stations = {};
stations["kone"] = new Set(["id", "nv", "ut"])
stations["ktwo"] = new Set(["wa", "id", "mt"])
stations["kthree"] = new Set(["or", "nv", "ca"])
stations["kfour"] = new Set(["nv", "ut"])
stations["kfive"] = new Set(["ca", "az"])

const finalStations = new Set()

// loops over our list of states needed and assembles a greedy list of all the
// stations that cover the most states
while ( statesNeeded.size > 0 ) {
    // best station and the states it covers
    let bestStation = undefined;
    let statesCovered = new Set();

    // loop over each station
    for (let station in stations) {
        // states this station covers
        const states = stations[station];
        // all unique states this station covered that aren't being covered by another station in {finalStations}
        const covered = new Set([...statesNeeded].filter(x => states.has(x)));

        // if the current covered set of stations is better than our current best covered in this loop
        if (covered.size > statesCovered.size) {
            // update our best stations and the states it covers
            bestStation = station;
            statesCovered = covered;
        }
    }

    // remove the states our latest best covered
    statesNeeded = new Set([...statesNeeded].filter(x => !statesCovered.has(x) ));
    // and add that station to the list
    finalStations.add(bestStation)
}