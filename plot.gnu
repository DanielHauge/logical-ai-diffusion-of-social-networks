# GNUPLOT script

# Set data file and output file
datafile = "gg.txt"
outputfile = "output.png"

# Set plot style and labels
set style data linespoints
set key top right
set title "Diffusion of AI involvement over a uniformly susceptible social network"
# Set subtitle
set xlabel "Transformations (k)"
set ylabel "Agents"
# Set legend to the right of the plot

# Plot the data
plot datafile using 0:2 with linespoints title 'Involved', \
     datafile using 0:3 with linespoints title 'Not Involved', \
     datafile using 0:4 with linespoints title 'Acted', \
     datafile using 0:5 with linespoints title 'Not Acted', \
     datafile using 0:6 with linespoints title 'Undecided'

# Uncomment the following lines if you have more columns to plot
# plot datafile using 5:6 with linespoints title 'Column 5 vs Column 6', \
#      datafile using 7:8 with linespoints title 'Column 7 vs Column 8'

# Save the plot to a file
set terminal pngcairo enhanced
set output outputfile
replot
set output

# Display the plot
# Uncomment the following line if you want to display the plot in a window
# pause -1
