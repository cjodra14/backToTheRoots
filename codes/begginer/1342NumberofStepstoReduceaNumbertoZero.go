package begginer

func NumberOfSteps(num int) int {
    var stepCount int
    for num != 0{
        stepCount++
        if num%2 == 0{
            num = num / 2

            continue
        } 

        num--
    }

    return stepCount 
}