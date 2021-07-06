<?php
if(php_sapi_name() != 'cli') die("Command Line Interface Required");

echo "Welcome to the quiz! Enter \"q\" to quit.\n";
$input = '';
while(!quit($input)){
    $question = new Question();
    $input = ask($question);
    answer($question, $input);
}
echo "Goodbye!\n";
sleep(2);


function quit(string $input): bool {
    return boolval(preg_match('/q|quit|stop|exit|end/i',$input));
}
class Question {
    private int $valueA = 0;
    private int $valueB = 0;
    private string $operator = '';
    public function __construct() {
        $this->valueA = rand(0,12);
        $this->valueB = rand(0,12);
        $this->operator = ['+','-','*'][rand(0,2)];
    }
    public function getFormula(): string {
        return "{$this->valueA} {$this->operator} {$this->valueB}";
    }
    public function getAnswer(): int {
        if($this->operator == '+')
            return $this->valueA + $this->valueB;
        if($this->operator == '-')
            return $this->valueA - $this->valueB;
        return $this->valueA * $this->valueB;
    }
}
function ask(Question $question): string {
    echo "What is ".$question->getFormula()."? ";
    return fread(STDIN,1024);
}
function answer(Question $question, string $input): void {
    if(quit($input)) return;
    $answer = $question->getAnswer();
    if($answer === intval($input)){
        echo "Correct! ".$question->getFormula()." = $answer!\n";
    } else {
        echo "WRONG :( ".$question->getFormula()." = $answer\n";
    }
}
