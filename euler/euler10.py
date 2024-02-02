import math

def euler1(n):
    sum = 0
    for num in range(n):
        if num%3 == 0 or num%5 == 0:
            sum += num
    return sum

def euler2a(n):
    if n <= 1:
        return n
    else:
        return euler2a(n-1) + euler2a(n-2)

def euler2b(n):
    sum = 0;
    for i in range(2,n+1):
        temp = euler2a(i);
        if temp%2 == 0:
            sum += temp
    return sum

def euler3(n):
    factor_arr = []
    x = 2
    while n > 1:
        while n%x == 0:
            factor_arr.append(int(x))
            n /= x
        x += 1
        if x*x > n:
            if n > 1: 
                factor_arr.append(int(n))
            break
    return max(factor_arr)

def euler4a(n):
    #n = str(n)
    #c = 0
    #state = True

    #while state == True:
    #    while c <= len(n)/2:
    #        if int(n[c]) == int(n[len(n) -c -1]):
    #            state = True
    #            c += 1
    #        else:
    #            state = False
    #            break
    #    break

    #if state == True:
    #    return True
    #else:
    #    return False

    #alternate way to detect palindrome, lol
    #gonna use this for memory purposes
    #"python doesnt care about memory!" "true"
    n = str(n)
    if n == n[::-1]:
        return True
    else:
        return False

def euler4b(n):
    #going in reverse is a pretty handy trick
    upper, max_prod = (10**(n))-1, 0
    lower = 1 + upper//10
    for x in range(upper,lower-1, -1):
        for y in range(x,lower-1,-1):
            prod = x * y
            if (prod < max_prod):
                break
            if (euler4a(prod) == True and prod > max_prod):
                max_prod = prod
         
    return max_prod


def euler5(n):
    count = 2
    x = 0
    while 1 > 0:
        for i in range(1, n+1):
            if count % i == 0:
                x += 1
            else:
                count += 1
                x = 0
        if x >= n-1:
            break
    return count

def euler6(n):
    arr = list(range(1,n+1))
    return abs(sum(arr)**2 - sum(list(map(lambda x: x**2, arr))))


    #same as:
    #sum_of_squares = sum(list(map(lambda x: x**2, arr)))
    #sum_of_natural_squared = sum(arr)**2
    #return abs(sum_of_squares - sum_of_natural_squared)

def euler7a(n):
    #detect if prime
    if n <= 1:
        return False
    for i in range(2, int(math.sqrt(n)+1)):
        if n % i == 0:
            return False
    return True

def euler7b(n):
    primes = []
    x = 2
    while len(primes) < n:
        if euler7a(x) == True:
            primes.append(x)
        x += 1;

    return primes[-1]

#def euler8():
#    big_num = "73167176531330624919225119674426574742355349194934
#96983520312774506326239578318016984801869478851843
#85861560789112949495459501737958331952853208805511
#12540698747158523863050715693290963295227443043557
#66896648950445244523161731856403098711121722383113
#62229893423380308135336276614282806444486645238749
#30358907296290491560440772390713810515859307960866
#70172427121883998797908792274921901699720888093776
#65727333001053367881220235421809751254540594752243
#52584907711670556013604839586446706324415722155397
#53697817977846174064955149290862569321978468622482
#83972241375657056057490261407972968652414535100474
#82166370484403199890008895243450658541227588666881
#16427171479924442928230863465674813919123162824586
#17866458359124566529476545682848912883142607690042
#24219022671055626321111109370544217506941658960408
#07198403850962455444362981230987879927244284909188
#84580156166097919133875499200524063689912560717606
#05886116467109405077541002256983155200055935729725
#71636269561882670428252483600823257530420752963450"

#def euler9()

def euler10():
    n = 2
    sumn = 0
    while n < 2000000:
        if euler7a(n) == True:
            sumn += n
        n += 1

    return sumn

if __name__ == "__main__":
    #print(euler1(1000))
    #print(euler2b(33))
    #print(euler3(600851475143))
    #print(euler4b(3))
    #print(euler5(20)) takes a bit
    #print(euler6(100))
    #print(euler7b(10001))
    print(euler10())

