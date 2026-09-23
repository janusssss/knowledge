BEGIN{
    count = 0
    dashes = sprintf("%*s", 143, "");
    gsub(/ /, "-", dashes);
    print dashes;
    printf "|%-38s|%-38s|%-38s|%-11s|%-12s|\n","volume_id", "attachment_id", "instance_uuid", "boot_index","multiattach"
}

NR != 1{
    if ($1 == prev) {
        printf "|%-38s|%-38s|%-38s|%-11s|%-12s|\n", "", $2, $3, $4,m;
    } else {
        count++
        print dashes;
        cmd = "cinder show " $1 " | grep multiattach " " | cut -c 35-40";
        cmd | getline m;
        close(cmd);
        printf "|%-38s|%-38s|%-38s|%-11s|%-12s|\n", $1, $2, $3, $4,m;
    }
    prev = $1
}

END{
    print dashes
    printf "存在多个attachment_id的卷数量: %s\n",count
}