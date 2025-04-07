library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.numeric_std.all;

entity andg2 is

    port (
        i_A : in  std_logic;            
-- input 1 to the AND gate
i_B : in  std_logic;            
            -- input 2 to the AND gate
            o_F : out std_logic             -- output of the AND gate
        );

end andg2;

architecture dataflow of andg2 is
begin
    o_F <= i_A and i_B;
end dataflow;
