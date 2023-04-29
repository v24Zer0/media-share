package com.mediashare.authenticationservice.util;

import java.util.regex.Pattern;

public class UUIDRegex {
    private final Pattern pattern;

    public UUIDRegex() {
        pattern = Pattern.compile("^[0-9a-fA-F]{8}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{12}$");
    }

    public boolean validateID(String id) {
        return this.pattern.matcher(id).matches();
    }
}
