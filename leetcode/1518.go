func numWaterBottles(numBottles int, numExchange int) int {
    emptyBottles := numBottles
    for emptyBottles >= numExchange {
        exchangedBottles := emptyBottles/numExchange
        numBottles += exchangedBottles
        emptyBottles -= (exchangedBottles*numExchange) - exchangedBottles
    }

    return numBottles
}
