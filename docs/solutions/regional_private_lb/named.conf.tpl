
options {
        listen-on port 53 { any; };
        allow-query    { localhost; ${vcn_cidr}; ${onprem_cidr}; };
        forward        only;
        forwarders     { 169.254.169.254; };
        recursion      yes;
};

zone "${onprem_domain}" {
        type       forward;
        forward    only;
        forwarders { ${onprem_dns_server1}; ${onprem_dns_server2}; };
};

zone "${zone}" {
        type  master;
        file  "/etc/named/db.${zone}";
};

