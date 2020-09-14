+++
# AUTOGENERATED BY byexample/generate.go
title= "Projects"
draft= false
description= ""
layout= "byexample"
weight = 4
topic = "Basics"
PlaygroundURL = "http://anz-bank.github.io/sysl-playground/?input=TW9iaWxlQXBwOgogICAgTG9naW46CiAgICAgICAgU2VydmVyIDwtIExvZ2luCiAgICAhdHlwZSBMb2dpbkRhdGE6CiAgICAgICAgdXNlcm5hbWUgPDogc3RyaW5nCiAgICAgICAgcGFzc3dvcmQgPDogc3RyaW5nCiAgICAhdHlwZSBMb2dpblJlc3BvbnNlOgogICAgICAgIG1lc3NhZ2UgPDogc3RyaW5nClNlcnZlcjoKICAgIExvZ2luKGRhdGEgPDogTW9iaWxlQXBwLkxvZ2luRGF0YSk6CiAgICAgICAgcmV0dXJuIE1vYmlsZUFwcC5Mb2dpblJlc3BvbnNlCg==&cmd="
GitRepoURL = "https://github.com/anz-bank/sysl/tree/master/demo/examples/Projects"
ID = "projects"
CodeWithoutComments = """MobileApp:
    Login:
        Server <- Login
    !type LoginData:
        username <: string
        password <: string
    !type LoginResponse:
        message <: string
Server:
    Login(data <: MobileApp.LoginData):
        return MobileApp.LoginResponse
"""

Segs = [[
  
      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>In this example we will use the &ldquo;call&rdquo; syntax to link two applications together.</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= true,CodeRendered="""<pre class="chroma">
<span class="nx">MobileApp</span><span class="p">:</span>
    <span class="nx">Login</span><span class="p">:</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">
        <span class="nx">Server</span> <span class="o">&lt;-</span> <span class="nx">Login</span></pre>""",DocsRendered= """<p>Here we specify that this endpoint has a dependency that it calls internally.</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="p">!</span><span class="kd">type</span> <span class="nx">LoginData</span><span class="p">:</span>
        <span class="nx">username</span> <span class="p">&lt;:</span> <span class="kt">string</span>
        <span class="nx">password</span> <span class="p">&lt;:</span> <span class="kt">string</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="p">!</span><span class="kd">type</span> <span class="nx">LoginResponse</span><span class="p">:</span>
        <span class="nx">message</span> <span class="p">&lt;:</span> <span class="kt">string</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">Server</span><span class="p">:</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">
    <span class="nf">Login</span><span class="p">(</span><span class="nx">data</span> <span class="p">&lt;:</span> <span class="nx">MobileApp</span><span class="p">.</span><span class="nx">LoginData</span><span class="p">):</span>
        <span class="k">return</span> <span class="nx">MobileApp</span><span class="p">.</span><span class="nx">LoginResponse</span></pre>""",DocsRendered= """<p>Use &ldquo;Application.Type&rdquo; to use a data type defined in another application</p>
""",Image = ""},

      {CodeEmpty= true,CodeLeading= false,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>TODO: Imports</p>
""",Image = ""},


],

]
+++

