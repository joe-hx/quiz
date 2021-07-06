async function main(){
    await printLn('Welcome to the quiz! Enter "q" to quit.')
    let input = ''
    while(!quit(input)){
        const q = new Question()
        input = await ask(q)
        await answer(q, input)
    }
    await printLn('Goodbye!')
    await later(2000)
}

function quit(input){
    return /q|quit|stop|exit|end/i.test(input)
}

class Question {
    #valueA = 0;
    #valueB = 0;
    #operator = ''
    constructor(){
        this.#valueA = Math.floor(Math.random() * 13)
        this.#valueB = Math.floor(Math.random() * 13)
        this.#operator = ['+','-','*'][Math.floor(Math.random() * 3)]
    }
    getFormula(){
        return `${this.#valueA} ${this.#operator} ${this.#valueB}`
    }
    getAnswer(){
        if(this.#operator === '+')
            return this.#valueA + this.#valueB
        if(this.#operator === '-')
            return this.#valueA - this.#valueB
        return this.#valueA * this.#valueB
    }
}

async function ask(q) {
    await print("What is "+q.getFormula()+"? ")
    return await readLn()
}

async function answer(q, a) {
    if(quit(a)) return;
    const answer = q.getAnswer()
    if(answer === parseInt(a)){
        await printLn(`Correct! ${q.getFormula()} = ${answer}!`)
    } else {
        await printLn(`WRONG :( ${q.getFormula()} = ${answer}`)
    }
}

async function print(str){
    await Deno.stdout.write(new TextEncoder().encode(str))
}

async function printLn(str){
    await print(''+str+'\n')
}

async function readLn(){
    const buf = new Uint8Array(1024)
    const n = await Deno.stdin.readSync(buf)
    return new TextDecoder().decode(buf.subarray(0, n)).trim()
}

function later(ms){
    return new Promise(function(resolve){
        setTimeout(function(){
            resolve()
        },ms)
    })
}

await main();
