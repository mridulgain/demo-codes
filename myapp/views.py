from django.http import HttpResponse
from django.shortcuts import render
from django.shortcuts import redirect

def home(request):
    return redirect('/blog')

def page1(request):
    return render(request, "index.html", {'author': "Mridul"})

def about(request):
    html_string = '''
    <h1> About page</h1>
    '''
    return HttpResponse(html_string)

def contacts(request):
    return render(request, "contacts.html", {'email': "Mridul@email.com"})

def error_404(request, exception):
    return render(request, 'error_404.html')

