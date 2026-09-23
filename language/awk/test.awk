BEGIN{
    Count = 0
}

{Count++}

{Count++}

END{
    print Count
}