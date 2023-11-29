datafile = "data.txt"
outputfile = "output.png"

set style data linespoints
set key top right
set title "Diffusion of AI involvement over a uniformly susceptible social network"
set xlabel "Transformations (k)"
set ylabel "Agents"

plot datafile using 0:2 with linespoints title 'Involved', \
     datafile using 0:3 with linespoints title 'Not Involved', \
     datafile using 0:4 with linespoints title 'Acted', \
     datafile using 0:5 with linespoints title 'Not Acted', \
     datafile using 0:6 with linespoints title 'Undecided'

set terminal pngcairo enhanced
set output outputfile
replot
set output

