function removeDuplicates(nums: number[]): number {
    if (nums.length === 0) return 0;

    let result = 0;

    for (let i = 1; i < nums.length; i++) {
        if (nums[result] !== nums[i]) {
            nums[++result] = nums[i]
        }
        console.log(nums);
    }

    return result + 1;
};

(function() {
  const nums = [0,0,1,1,1,2,2,3,3,4];
  const expected = [0,1,2,3,4];

  const k = removeDuplicates(nums);
  let success = false;

  console.log(`expected: ${expected} | result: ${nums.slice(0, k)}`)
}());
