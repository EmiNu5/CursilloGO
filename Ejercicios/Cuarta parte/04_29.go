package main

import "fmt"

func main29() {
	var txt string = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas sit amet leo vitae magna rutrum vestibulum. Aliquam erat volutpat. Praesent id nulla et orci gravida mollis. Vestibulum tincidunt interdum nunc sit amet malesuada. Pellentesque aliquam risus in magna ornare, vel dignissim lectus maximus. Sed euismod mi nec est laoreet consequat. Nam et velit quis dolor mollis ullamcorper.
Integer sodales ornare risus, at vehicula enim dapibus eget. Quisque at interdum turpis. Cras interdum suscipit magna sit amet congue. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In eu erat vel ante euismod tincidunt vel et nunc. Mauris varius scelerisque volutpat. Ut facilisis ipsum vitae hendrerit convallis. Proin tristique imperdiet dolor, eu laoreet massa suscipit non. Vestibulum lobortis, massa condimentum aliquet pulvinar, enim lectus imperdiet massa, placerat porta dui dolor ac mauris. Donec suscipit elit sem, id ultrices ligula placerat id. Vivamus dictum justo velit, eget tempus metus dapibus non. Sed venenatis est non lacus dapibus, posuere rutrum ipsum condimentum. Integer quis tincidunt libero, semper suscipit enim. Sed vel fermentum quam, et sodales dolor.
Mauris volutpat eros quis pretium sodales. Nulla ultricies mollis dui eget suscipit. Aenean pellentesque enim turpis, nec consectetur dolor pellentesque id. Vivamus in mi bibendum orci bibendum sagittis. Quisque dignissim imperdiet nibh, ut dignissim eros egestas vitae. Phasellus gravida vulputate congue. Nullam id rutrum lacus, nec tempor ipsum.
Quisque eu gravida enim. Quisque auctor neque ante, ac varius mi suscipit id. Integer nulla ex, accumsan vitae augue ac, faucibus iaculis elit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Fusce condimentum erat eu tortor laoreet vulputate sed sed massa. Morbi fermentum elit vitae elit semper elementum at vulputate enim. Aenean id nisl mollis, pulvinar est sit amet, pretium diam.
Morbi nec leo vel massa hendrerit euismod eu ut nisl. Sed dignissim, risus et interdum interdum, eros augue viverra urna, a sodales sapien turpis a justo. Donec quis tortor ipsum. Praesent tempor sit amet purus non auctor. Mauris et pharetra quam. Ut sed laoreet nulla. Ut dictum aliquet neque, malesuada finibus enim imperdiet nec. Sed id ultrices mi. Sed sit amet venenatis turpis. Aliquam erat volutpat.	
`
	const enter byte = '\n'
	const space byte = ' '
	const ipsum string = "ipsum"
	const quisque string = "Quisque"
	const spaceString string = " "

	fmt.Println("Cantidad de Parrafos: ", len(findAllChar(txt, enter)))

	fmt.Println("Cantidad de Palabras: ", len(findAllChar(txt, space))+len(findAllChar(txt, enter)))

	fmt.Println("Cantidad de Caracteres: ", len(txt)-len(findAllChar(txt, enter)))

	fmt.Println("Cantidad de Ipsum: ", len(findAllString(txt, ipsum)))

	fmt.Println("Indice de la primer palabra: ", findFirstString(txt, quisque))

	fmt.Println("El reverso es: \n", join(reverseString(splitByString(txt, spaceString)), spaceString))
}
