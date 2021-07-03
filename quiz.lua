math.randomseed(os.time())
print 'Welcome to the quiz! Enter "q" to quit'

function quit(str)
    str = tostring(str)
    return str == 'q' or str == 'quit' or str == 'end' or str == 'exit' or str == 'stop'
end

local input = ''
local ops = {'+','-','*'}

while not quit(input) do
    local valueA = math.random(0,12)
    local valueB = math.random(0,12)
    local op = math.random(1, rawlen(ops));
    local operator = ops[op];

    io.write('What is '..tostring(valueA)..' '..operator..' '..tostring(valueB)..'? ')
    input = io.read()

    if not quit(input) then
        local given = tonumber(input)

        local answer = 0
        if operator == '+' then
            answer = valueA + valueB
        elseif operator == '-' then
            answer = valueA - valueB
        else
            answer = valueA * valueB
        end

        if(answer == given) then
            print('Correct! '..tostring(valueA)..' '..operator..' '..tostring(valueB)..' = '..tostring(answer)..'!')
        else
            print('Wrong :( '..tostring(valueA)..' '..operator..' '..tostring(valueB)..' = '..tostring(answer))
        end
    end

end
print 'Goodbye!'


