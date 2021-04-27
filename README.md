Nia Therapeutics Interview question

PROMPT:

You are developing a software application. One of the requirements is to display a splash screen while your application is starting up. To make time go by faster for the user, you need to implement a fade-in effect for the splash screen: starting from a blank screen, a splash screen image should emerge pixel by pixel over a span of several seconds until the entire image is shown. Source pixel location at each iteration must be chosen randomly (or pseudo randomly), so the image transfer does not form any discernible pattern.

Your task is to devise and implement the most efficient algorithm which takes image dimensions as its input, and outputs a randomized sequence of pixel coordinates for pixel-by-pixel image transfer to the screen. That is, for an image of width W and height H, your algorithm should produce a sequence of distinct pixel coordinates with W*H elements. Your implementation should be in Python, C, C++ or Go and utilize only the basic language constructs such as variables, arrays, lists, conditionals and loops. "rand" function or its equivalent can be used for random number generation. Please accompany your implementation with a theoretical performance analysis of your algorithm in terms of computational complexity and memory requirements. Compare your final solution with other less efficient approaches that you have considered.

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

RESPONSE:

For computational complexity, I am considering n = W*H, which is equal to the length of the response sequence in every case.
Thus the algorithm I've devised runs in O(n) time:

1. Create a list of length W*H of vertice names                         O(W*H)
2. Create an index of length W*H                                        O(W*H)
3. Shuffle index in place (Fisher-Yates method)                         O(W*H)
4. Print each vertex array element in the order of the shuffled index   O(W*H)
