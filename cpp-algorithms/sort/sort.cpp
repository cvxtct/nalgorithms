#include <iostream>
#include <vector>


int get_next_gap(int gap);
std::vector<double> comb_sort(std::vector<double> arr);

int get_next_gap(int gap) {
    gap = (gap * 10) / 13;
    if(gap < 1) {
        return 1;
    }
    return gap;
};


std::vector<double> comb_sort(std::vector<double> arr) {
    int n = arr.size();
    int gap = n;
    bool swapped = true;
    int stepps = 0;
    while (swapped || gap != 1){
        gap = get_next_gap(gap);
        swapped = false;
        for (size_t i{0}; i < n-gap; i++){
            if(arr[i] > arr[i+gap]){
                arr[i], arr[i+gap] = arr[i+gap], arr[i];
                stepps++;
                swapped = true;
            }
        }
    }
    return arr;
}


int main() {

}


