/*
 * Compiled with go version go1.9 linux/amd64
 */
package main

import (
	"fmt"
	"math"
	"os"
)

/* Following function computes
 * means of a given data set.
 */
func compute_mean(data_points []float32) float32 {
	var (
		i                  int
		sum                float32
		number_of_elements int
	)
	number_of_elements = len(data_points)
	for i = 0; i < number_of_elements; i++ {
		sum = sum + data_points[i]
	}
	return (sum / float32(number_of_elements))
}

/* Following function computes standard
 * deviation of a data point
 */
func compute_std_deviation(data_points []float32) float32 {
	var (
		number_of_elements int
		i                  int
		mean               float32
		data_point_dev     []float32
		std_div, temp      float32
	)
	number_of_elements = len(data_points)
	data_point_dev = make([]float32, number_of_elements)
	/*Calculate mean of the data points*/
	mean = compute_mean(data_points)
	/*Calculate standard deviation of individual data points*/
	for i = 0; i < number_of_elements; i++ {
		temp = data_points[i] - mean
		data_point_dev[i] = (temp * temp)
	}

	/* Finally, Compute standard Deviation of data points
	 * by taking mean of individual std. Deviation.
	 */
	std_div = compute_mean(data_point_dev)
	return float32(math.Sqrt(float64(std_div)))
}

/* Following function computes covariance
 * of two data sets.
 */
func covariance(data_point_x []float32, data_point_y []float32) float32 {
	var (
		i      int
		len_x  int
		mean_x float32
		mean_y float32
		sum    float32
	)
	/*Check number of data points*/
	len_x = len(data_point_x)
	if len_x != len(data_point_y) {
		fmt.Println("Both data set must of equal size!")
		os.Exit(1)
	}

	/*Calculate mean of data points*/
	mean_x = compute_mean(data_point_x)
	mean_y = compute_mean(data_point_y)
	/*Finally, covariance*/
	for i = 0; i < len_x; i++ {
		sum = sum + ((data_point_x[i] - mean_x) * (data_point_y[i] - mean_y))
	}
	return (sum / float32(len_x))
}

/* Following function calculates the correlation
 * coefficient of given data sets.
 */
func compute_corelation_coeffecient(data_point_x []float32, data_point_y []float32) float32 {

	var (
		covariance_x_y, std_div_x, std_div_y float32
	)
	covariance_x_y = covariance(data_point_x, data_point_y)
	std_div_x = compute_std_deviation(data_point_x)
	std_div_y = compute_std_deviation(data_point_y)
	return (covariance_x_y / (std_div_x * std_div_y))
}

/* Function computes Manhattan distance between two data sets */
func compute_manhattan_distance(data_point_x, data_point_y []float32) float32 {
	var i int
	var sum float32
	if len(data_point_x) != len(data_point_y) {
		fmt.Println("Both data sets must be equal in numbers!")
		os.Exit(1)
	}
	for i = 0; i < len(data_point_x); i++ {
		sum = sum + float32(math.Abs(float64(data_point_x[i])-float64(data_point_y[i])))
	}
	return sum
}

/* Function Calculates Euclidean distance
 * between two data sets
 */
func compute_Euclidean_distance(data_point_x, data_point_y []float32) float32 {
	var i int
	var sum float32
	if len(data_point_x) != len(data_point_y) {
		fmt.Println("Both data sets must be equal in numbers!")
		os.Exit(1)
	}
	for i = 0; i < len(data_point_x); i++ {
		sum = sum + ((data_point_x[i] - data_point_y[i]) * (data_point_x[i] - data_point_y[i]))
	}
	return float32(math.Sqrt(float64(sum)))
}

/* Compute Chebyshev distance between two data points */
func compute_Chebyshev_distance(data_point_x, data_point_y []float32) float32 {
	var (
		distance float32
		i        int
	)
	if len(data_point_x) != len(data_point_y) {
		fmt.Println("Data points must be equal!")
		os.Exit(1)
	}
	distance = 0
	for i = 0; i < len(data_point_y); i++ {
		if distance < float32(math.Abs(float64(data_point_x[i]-data_point_y[i]))) {
			distance = float32(math.Abs(float64(data_point_x[i] - data_point_y[i])))
		}
	}
	return distance
}

func main() {
	data_point_a := []float32{2, 4, 4, 4, 5, 5, 7, 9}
	data_point_b := []float32{2, 4, 4, 4, 5, 5, 7, 1}
	fmt.Println("Standard Deviation:", compute_std_deviation(data_point_a))
	fmt.Println("Chebyshev Distance:", compute_Chebyshev_distance(data_point_a, data_point_b))
	fmt.Println("Correlation Coeffiecent:", compute_corelation_coeffecient(data_point_a, data_point_b))
}
