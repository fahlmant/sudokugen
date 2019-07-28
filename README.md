# Sudokugen

WIP CLI and Go pkg for generating arbitrarily sized square Sudoku solutions. 

The CLI takes two arguments: the x dimesion and y dimension of a single box of the puzzle.
If the desired puzzle is a 9x9, then the command would be `sudokugen 3 3`. For a 12x12, `sudokugen 4 3` or `sudokugen 3 4`.
Only square puzzles can be created. 

TODO:
--json
--pretty
--write