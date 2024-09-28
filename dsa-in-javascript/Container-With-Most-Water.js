function maxArea(arr)
{
    let left = 0;
    let right = arr.length-1;
    let maxarea = 0;

    while(left < right){
        let ht =  Math.min(arr[left], arr[right]);
         
        maxarea = Math.max(maxarea, ht * (right-left));
       

        if(arr[left] < arr[right]){
            left++;
        }else{
            right--;
        }
    }
     
    return maxarea;

}

const arr = [1,8,6,2,5,4,8,3,7];
console.log(maxArea(arr));