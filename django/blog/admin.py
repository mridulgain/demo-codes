from django.contrib import admin
from .models import Comment, Post
# Register your models here.
# admin.site.register(Comment)

@admin.register(Post)
class PostAdmin(admin.ModelAdmin):
    list_display = ('title', 'slug', 'author', 'publish', 'status')
    prepopulated_fields = {'slug': ("title", )}
    date_hierarchy = 'publish'
    ordering = ('status', 'publish')
    list_filter = ('status', 'created', 'publish', 'author')
    search_fields = ('title', 'body')
    # raw_id_fields = ('author',)

@admin.register(Comment)
class CommentAdmin(admin.ModelAdmin):
    list_display = ('name', 'email', 'post', 'created', 'active')
    list_filter = ('active', 'created', 'updated')
    search_fields = ('name', 'email', 'body')
