from plyer import notification
#from pyramid.compat import text_, string_types


title = "Hello Amazing people!"

message = "Thankyou for reading! Take care!"

notification.notify(title=title,
                    message=message,
                    app_icon=None,
                    timeout=10,
                    toast=False)
