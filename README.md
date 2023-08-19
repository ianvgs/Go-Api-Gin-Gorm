You can check this project running at aws-ec2-ubuntu machine here: http://3.93.242.129.

See you, yours Ian

### Run server  changes
/gobuild
./goagain

### Run Local + watching changes
CompileDaemon -command="./your-app-binary"
CompileDaemon -command="./goagain"  

Logging:

No html:
A) Direto no html => 
             <h4>
                {{ $element.Titulo}}
            </h4>
            ***Lembrar que o atrituboto é renomeado na model titulo => Titulo;

B) No browser:
            <script>
                var titulo = "{{ $element.Titulo }}";
                console.log("Titulo:", titulo);
            </script>            
            <script>
                    var jsonData = {{ .element }};
                    console.log("JSON Data:", jsonData);
            </script>  
            <script>
                    var jsonData = {{ .news }};
                    console.log("JSON Data:", jsonData);
            </script>
             ***Pode inserir dentro de qualquer tag





No terminal (arquivos fora do html):
    Log Unico:
     fmt.Println("Noticias:", noticias) 
	

    Log de array:    
     for _, noticia := range noticias {
		fmt.Println("Colaborador:", noticia.Colaborador)
	} 