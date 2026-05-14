:local fileName "iran.rsc"
:local url "https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc"

/tool fetch url=$url dst-path=$fileName mode=http

:if ([:len [/file find name=$fileName]] = 0) do={
    :log error "Fetch failed - file not found"
    :return
}

:if ([/file get $fileName size] < 10) do={
    :log error "File too small - abort"
    /file remove $fileName
    :return
}

:log info "Importing $fileName"
/import file-name=$fileName
:log info "Import done"

/file remove $fileName
