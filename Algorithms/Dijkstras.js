const graph = {}

graph["start"] = {}
graph["start"]["a"] = 6;
graph["start"]["b"] = 2;

graph["a"] = {};
graph["a"]["fin"] = 1;

graph["b"] = {};
graph["b"]["a"] = 3;
graph["b"]["fin"] = 5;

graph["fin"] = {}

const costs = {}
costs["a"] = 6;
costs["b"] = 2;
costs["fin"] = Infinity;

const parents = {};
parents["a"] = "start";
parents["b"] = "start";
parents["fin"] = undefined;

const processed = [];

function findLowestCostNode() {
    let smallestCost = Infinity;
    let smallestKey = undefined;

    for (let key in costs) {
        if (processed.indexOf(key) === -1 && costs[key] < smallestCost) {
            smallestKey = key
        }
    }

    return smallestKey;
}

let node = findLowestCostNode()
while (node !== undefined) {
    const cost = costs[node];
    const neighbors = graph[node];

    for (let n in neighbors) {
        const newCost = cost + neighbors[n];

        if (costs[n] > newCost) {
            costs[n] = newCost;
            parents[n] = node;
        }
    }

    processed.push(node);
    node = findLowestCostNode();
}

const path = ["fin"];

let pathNode = parents["fin"];
while (pathNode !== undefined) {
    path.push(pathNode)
    pathNode = parents[pathNode]
}

console.log(path)