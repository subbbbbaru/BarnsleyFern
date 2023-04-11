# BarnsleyFern
The Barnsley fern is a fractal named after the British mathematician Michael Barnsley who first described it in his book Fractals Everywhere.
He made it to resemble the black spleenwort, Asplenium adiantum-nigrum. 
The fern is one of the basic examples of self-similar sets, i.e. it is a mathematically generated pattern that can be reproducible at any magnification or reduction. 
Like the Sierpinski triangle, the Barnsley fern shows how graphically beautiful structures can be built from repetitive uses of mathematical formulas with computers.
Barnsley's fern uses four affine transformations. The formula for one transformation is the following:

![7a6ccaaa702062c1c63c488c076d4e1d08f94302](https://user-images.githubusercontent.com/45457578/231210876-febbcea4-1213-43c9-9024-9d5d0489b9e4.jpg)

Barnsley shows the IFS code for his Black Spleenwort fern fractal as a matrix of values shown in a table.
In the table, the columns "a" through "f" are the coefficients of the equation, and "p" represents the probability factor.

| w     | a     | b     | c     | d     | e     | f     | p     | Portion generated |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |  :---: |
| ƒ1    | 0 	  | 0 	  |   0 	  | 0.16 |	0 |	0 	 | 0.01  |	Stem|
| ƒ2 | 0.85 	 | 0.04 |	−0.04 | 0.85 |	0 |	1.60 |	0.85 |	Successively smaller leaflets|
| ƒ3 | 0.20 	| −0.26 |	0.23 	| 0.22 |	0 |	1.60 |	0.07 |	Largest left-hand leaflet|
| ƒ4 | −0.15 |  0.28 |	0.26 | 0.24 |	0 |	0.44 |	0.07 |	Largest right-hand leaflet| 

By playing with the coefficients, it is possible to create mutant fern varieties. In his paper on V-variable fractals, Barnsley calls this trait a superfractal.


![fern](https://user-images.githubusercontent.com/45457578/231211067-c754c36c-9ad9-4d4b-a27f-a5672f4abf89.png)
![34](https://user-images.githubusercontent.com/45457578/231214841-44e56b16-2828-4a93-8310-c98011fa435b.png)
![34](https://user-images.githubusercontent.com/45457578/231215125-72a72e7e-c261-4f68-a42e-55ee5ff653cd.png)


